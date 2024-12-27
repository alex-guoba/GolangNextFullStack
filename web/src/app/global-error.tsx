"use client"; // Error boundaries must be Client Components

import {
  PageHeader,
  PageHeaderDescription,
  PageHeaderHeading,
} from "@/components/page-header";
import Shell from "@/components/shell/shell";
import { useEffect } from "react";
import { cn } from "@/lib/utils";
import { Button, buttonVariants } from "@/components/ui/button";

export default function GlobalError({
  error,
  reset,
}: {
  error: Error & { digest?: string };
  reset: () => void;
}) {
  useEffect(() => {
    // Log the error to an error reporting service
    console.error(error);
  }, [error]);

  return (
    <html>
      <body>
        <div>
          <Shell as="article" className="relative flex min-h-screen flex-col">
            <PageHeader>
              <PageHeaderHeading>Error</PageHeaderHeading>
              <PageHeaderDescription size="sm" className="text-center">
                Something went wrong, try again later
              </PageHeaderDescription>
            </PageHeader>

            <Button
              className={cn(
                buttonVariants({
                  variant: "ghost",
                  className: "mx-auto mt-4 w-fit",
                })
              )}
              onClick={() => reset()}
            >
              Try again
            </Button>
          </Shell>
        </div>
      </body>
    </html>
  );
}
