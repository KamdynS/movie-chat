# Clerk Integration Guide

## 1. Frontend Integration

1. Install Clerk SDK:
   ```
   npm install @clerk/nextjs
   ```

2. Update `_app.tsx` to wrap your application with ClerkProvider:
   ```tsx
   import { ClerkProvider } from '@clerk/nextjs';

   function MyApp({ Component, pageProps }) {
     return (
       <ClerkProvider>
         <Component {...pageProps} />
       </ClerkProvider>
     );
   }

   export default MyApp;
   ```

3. Replace existing auth components with Clerk components:
   - Use `<SignIn />` for sign-in
   - Use `<SignUp />` for sign-up
   - Use `<UserButton />` for user menu

4. Protect routes using Clerk's middleware. Create a `middleware.ts` file in your root directory:
   ```typescript
   import { authMiddleware } from "@clerk/nextjs";

   export default authMiddleware({
     publicRoutes: ["/", "/api/public"]
   });

   export const config = {
     matcher: ["/((?!.*\\..*|_next).*)", "/", "/(api|trpc)(.*)"],
   };
   ```

## 2. Backend Integration

1. Install Clerk's Go SDK:
   ```
   go get github.com/clerkinc/clerk-sdk-go
   ```

2. Set up Clerk client in your main Go file:
   ```go
   import (
     "github.com/clerkinc/clerk-sdk-go/clerk"
   )

   func main() {
     clerkClient, err := clerk.NewClient("your_clerk_secret_key")
     if err != nil {
       log.Fatal(err)
     }
     // Use clerkClient in your application
   }
   ```

3. Create a middleware to verify Clerk sessions:
   ```go
   func ClerkAuthMiddleware() gin.HandlerFunc {
     return func(c *gin.Context) {
       sessionToken := c.GetHeader("Authorization")
       if sessionToken == "" {
         c.JSON(http.StatusUnauthorized, gin.H{"error": "No authorization header provided"})
         c.Abort()
         return
       }

       claims, err := clerkClient.VerifyToken(sessionToken)
       if err != nil {
         c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid session token"})
         c.Abort()
         return
       }

       c.Set("userId", claims.Subject)
       c.Next()
     }
   }
   ```

4. Apply the middleware to your protected routes:
   ```go
   protected := r.Group("/")
   protected.Use(ClerkAuthMiddleware())
   {
     protected.POST("/create-room", createRoomHandler)
     // other protected routes
   }
   ```

## 3. Database Schema Update

Since Clerk will handle user authentication, you'll need to update your database schema to link Clerk's user IDs with your application's user data. Here's a suggested schema for a PostgreSQL table:

```sql
CREATE TABLE user_profiles (
    id SERIAL PRIMARY KEY,
    clerk_user_id VARCHAR(255) UNIQUE NOT NULL,
    username VARCHAR(255) UNIQUE,
    email VARCHAR(255) UNIQUE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Index for faster lookups
CREATE INDEX idx_clerk_user_id ON user_profiles(clerk_user_id);
```

This table will store additional user information that Clerk doesn't manage. The `clerk_user_id` field will be used to link your application's user data with Clerk's user records.

## 4. Update User-related Operations

1. On user sign-up or first login, create a record in the `user_profiles` table:
   ```go
   func createUserProfile(db *sql.DB, clerkUserId, username, email string) error {
     _, err := db.Exec(`
       INSERT INTO user_profiles (clerk_user_id, username, email)
       VALUES ($1, $2, $3)
       ON CONFLICT (clerk_user_id) DO NOTHING
     `, clerkUserId, username, email)
     return err
   }
   ```

2. Update your user-related queries to use the `clerk_user_id` instead of your previous user ID.

Remember to remove any existing authentication logic (like password hashing, JWT generation, etc.) as Clerk will handle these aspects of authentication.