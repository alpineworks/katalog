import { cn } from "@/lib/utils";
import { pacifico } from "./fonts";
import { Button } from "@/components/ui/button";
import Link from "next/link";

export default function NotFound() {
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
              not found!
            </h1>
            <Link href="/">
              <Button variant="outline" className="mt-4">
                go home
              </Button>
            </Link>
          </div>
        </div>
      </div>
    </main>
  );
}
