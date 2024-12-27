import { MobileNav } from "@/components/layouts/mobile-nav";
import { MainNav } from "@/components/layouts/main-nav";
import { siteConfig } from "@/config/site";
import Shell from "../shell/shell";

export async function SiteHeader() {
  return (
    <header className="sticky top-0 z-40 w-full border-b bg-background">
      <Shell as="nav" className="flex h-16 items-center">
        <MainNav items={siteConfig.mainNav} />
        <MobileNav mainNavItems={siteConfig.mainNav} />
      </Shell>
    </header>
  );
}
