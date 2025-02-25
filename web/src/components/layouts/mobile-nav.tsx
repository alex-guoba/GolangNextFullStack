"use client";

import * as React from "react";
import Link from "next/link";
import { useSelectedLayoutSegment } from "next/navigation";
import { siteMeta } from "@/config/meta";
import { cn } from "@/lib/utils";
import {
  Accordion,
  AccordionContent,
  AccordionItem,
  AccordionTrigger,
} from "@/components/ui/accordion";
import { Button } from "@/components/ui/button";
import { ScrollArea } from "@/components/ui/scroll-area";
import {
  Sheet,
  SheetContent,
  SheetHeader,
  SheetTitle,
  SheetTrigger,
} from "@/components/ui/sheet";
import { Icons } from "@/components/icons";
import { MainNavItem } from "@/config/site";

interface MobileNavProps {
  mainNavItems?: MainNavItem[];
}

export function MobileNav({ mainNavItems }: MobileNavProps) {
  const segment = useSelectedLayoutSegment();
  const [isOpen, setIsOpen] = React.useState(false);

  const navItems = mainNavItems || [];

  return (
    <Sheet open={isOpen} onOpenChange={setIsOpen}>
      <SheetTrigger asChild>
        <Button
          variant="ghost"
          size="icon"
          className="size-5 hover:bg-transparent focus-visible:bg-transparent focus-visible:ring-0 focus-visible:ring-offset-0 lg:hidden"
        >
          <Icons.menu aria-hidden="true" />
          <span className="sr-only">Toggle Menu</span>
        </Button>
      </SheetTrigger>
      <SheetContent side="left" className="pl-1 pr-0">
        <SheetHeader>
          <SheetTitle className="px-7">
            <Link
              href="/"
              className="flex items-center"
              onClick={() => setIsOpen(false)}
            >
              <Icons.logo className="mr-2 h-4 w-4" aria-hidden="true" />
              <span className="font-bold">{siteMeta.name}</span>
              <span className="sr-only">Home</span>
            </Link>
          </SheetTitle>
          {/* <SheetDescription>
            Make changes to your profile here. Click save when you're done.
          </SheetDescription> */}
        </SheetHeader>
        <ScrollArea className="my-4 h-[calc(100vh-8rem)] pb-10 pl-6">
          <div className="pl-1 pr-7">
            <Accordion
              type="multiple"
              defaultValue={navItems.map((item) => item.title!)}
              className="w-full"
            >
              {navItems?.map((item, index) =>
                item.items && item.items?.length > 0 ? (
                  <AccordionItem value={item.title!} key={index}>
                    <AccordionTrigger className="text-sm capitalize">
                      {item.title}
                    </AccordionTrigger>
                    <AccordionContent>
                      <div className="flex flex-col space-y-2">
                        {item.items?.map((subItem, index) =>
                          subItem.href ? (
                            <MobileLink
                              key={index}
                              href={String(subItem.href)}
                              segment={String(segment)}
                              setIsOpen={setIsOpen}
                              disabled={subItem.disabled}
                            >
                              {subItem.title}
                            </MobileLink>
                          ) : (
                            <div
                              key={index}
                              className="text-foreground/70 transition-colors"
                            >
                              {item.title}
                            </div>
                          )
                        )}
                      </div>
                    </AccordionContent>
                  </AccordionItem>
                ) : (
                  item.href && (
                    <AccordionItem key={index} value={item.title!}>
                      <Link href={item.href} legacyBehavior passHref>
                        {/* <MobileLink
                          className={cn(navigationMenuTriggerStyle(), "h-auto")}
                        > */}
                        <div className="flex flex-col space-y-2 py-4 font-medium transition-all hover:underline">
                          {item.title}
                        </div>
                        {/* </MobileLink> */}
                      </Link>
                    </AccordionItem>
                  )
                )
              )}
            </Accordion>
          </div>
        </ScrollArea>
      </SheetContent>
    </Sheet>
  );
}

interface MobileLinkProps extends React.PropsWithChildren {
  href: string;
  disabled?: boolean;
  segment: string;
  setIsOpen: React.Dispatch<React.SetStateAction<boolean>>;
}

function MobileLink({
  children,
  href,
  disabled,
  segment,
  setIsOpen,
}: MobileLinkProps) {
  return (
    <Link
      href={href}
      className={cn(
        "text-foreground/70 transition-colors hover:text-foreground",
        href.includes(segment) && "text-foreground",
        disabled && "pointer-events-none opacity-60"
      )}
      onClick={() => setIsOpen(false)}
    >
      {children}
    </Link>
  );
}
