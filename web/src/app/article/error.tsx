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

export default function Error({
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
    <Shell variant="centered" className="relative flex min-h-[80vh] flex-col">
      <PageHeader className="text-center">
        <PageHeaderHeading>Page Error</PageHeaderHeading>
        <PageHeaderDescription size="sm" className="text-center">
          Something went wrong, try again later
        </PageHeaderDescription>
      </PageHeader>

      <Button
        className={cn(
          buttonVariants({
            variant: "ghost",
            className: "mx-auto mt-2 w-fit",
          })
        )}
        onClick={() => reset()}
      >
        Try again
      </Button>
      <span className="sr-only">Try again</span>
    </Shell>
  );
}
