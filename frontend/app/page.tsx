"use client";
import { pacifico } from "./fonts";
import { cn } from "@/lib/utils";

export default function Home() {
  return (
    <main>
      <div className="flex flex-col items-center justify-center min-h-[calc(100vh-80px)]">
        <div className="text-center">
          <div className="max-w-xl">
            <h1
              className={cn(
                "text-8xl p-16 bg-gradient-to-bl from-fuchsia-500 to-cyan-500 overflow-visible	inline-block text-transparent bg-clip-text",
                pacifico.className
              )}
            >
              katalog!
            </h1>
          </div>
        </div>
      </div>
    </main>
  );
}
