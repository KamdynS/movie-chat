import { UserProfile } from '@clerk/nextjs'

const UserProfilePage = () => (
  <div className="container mx-auto py-8">
    <UserProfile path="/user-profile" />
  </div>
)

export default UserProfilePage
