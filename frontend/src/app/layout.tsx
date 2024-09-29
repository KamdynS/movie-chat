import type { Metadata } from "next";
import localFont from "next/font/local";
import "./globals.css";
import Navbar from "@/components/Navbar"
import {
  ClerkProvider,
  // Remove unused imports
} from '@clerk/nextjs'

const geistSans = localFont({
  src: "./fonts/GeistVF.woff",
  variable: "--font-geist-sans",
  weight: "100 900",
});
const geistMono = localFont({
  src: "./fonts/GeistMonoVF.woff",
  variable: "--font-geist-mono",
  weight: "100 900",
});

export const metadata: Metadata = {
  title: "Movie Chatrooms",
  description: "Chatrooms for discussing movies",
};

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="en">
      <ClerkProvider
        appearance={{
          variables: { colorPrimary: "#000000" },
          elements: {
            // ... existing elements ...
            // Add styling for user profile page
            userProfilePage: "bg-slate-900 text-slate-100",
            userProfile: "bg-slate-800 border-slate-700",
            userProfileSection: "bg-slate-800",
            userButtonBox: "bg-slate-700 hover:bg-slate-600",
          },
        }}>
        <body
          className={`${geistSans.variable} ${geistMono.variable} antialiased bg-slate-900`}
        >
          <Navbar />
          <main className="bg-slate-900 text-slate-100">
            {children}
          </main>
        </body>
      </ClerkProvider>
    </html>
  );
}