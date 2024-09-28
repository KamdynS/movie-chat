'use client';

import Link from 'next/link'
import { Button } from './ui/button'
import { useRouter } from 'next/navigation';
import { UserButton, SignInButton, SignUpButton, useUser } from "@clerk/nextjs";

export default function Navbar() {
  const { isSignedIn, user } = useUser();
  const router = useRouter();

  return (
    <header className="px-4 py-6 sm:px-6 lg:px-8 bg-slate-900">
      <div className="container max-w-5xl mx-auto flex items-center justify-between">
        <Link href="/" className="flex items-center gap-2">
          <span className="text-xl font-bold text-slate-100">Chatrooms for a movie</span>
        </Link>
        <nav className="flex items-center gap-4">
          <Link href="/" className="text-sm font-medium text-slate-100 hover:underline underline-offset-4">
            Home
          </Link>
          <Link href="/guide" className="text-sm font-medium text-slate-100 hover:underline underline-offset-4">
            Guide
          </Link>
          <Link href="/about" className="text-sm font-medium text-slate-100 hover:underline underline-offset-4">
            About
          </Link>
        </nav>
        <div className="flex items-center gap-2">
          {!isSignedIn ? (
            <>
              <SignInButton mode="modal">
                <Button variant="outline" className="bg-slate-100 text-purple-700 hover:bg-purple-700">Sign in</Button>
              </SignInButton>
              <SignUpButton mode="modal">
                <Button className="bg-purple-700 hover:bg-purple-600 text-slate-100">Sign up</Button>
              </SignUpButton>
            </>
          ) : (
            <>
              <Link href="/profile" className="text-sm font-medium text-slate-100 hover:underline underline-offset-4">
                Profile
              </Link>
              <UserButton afterSignOutUrl="/" />
            </>
          )}
        </div>
      </div>
    </header>
  )
}