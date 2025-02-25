import Shell from "@/components/shell/shell";
import { buttonVariants } from "@/components/ui/button";
import { cn } from "@/lib/utils";
import Link from "next/link";

// export const dynamic = "force-dynamic";

export default async function NotFound() {
  return (
    <>
      <Shell variant="centered" className="">
        <section className="my-4 text-center">
          <p className="text-[6rem] font-semibold">404</p>

          <h1 className="mt-4 text-3xl font-bold tracking-tight sm:text-5xl">
            Page not found
          </h1>

          <div className="mt-10 flex items-center justify-center gap-x-6">
            <Link
              href="/"
              className={cn(
                buttonVariants({ className: "mx-auto mt-4 w-fit" })
              )}
            >
              {" "}
              Go back home{" "}
            </Link>
          </div>
        </section>
      </Shell>
    </>
  );
}
