"use server"

import { signIn, signOut } from "@/lib/auth/auth";

export async function SignIn() {
    await signIn(undefined);
}

export async function SignInWithProvider(providerId: string, callbackUrl: string) {
    await signIn(providerId, { redirectTo: callbackUrl });
}

export async function SignInWithProviderDefaultCallback(providerId: string) {
    await signIn(providerId, { redirectTo: "/deployments" });
}

export async function SignOut() {
    await signOut( { redirectTo: "/" });
}