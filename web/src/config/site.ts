// import type { MainNavItem } from "@/types";

import type { Icons } from "@/components/icons";

export interface NavItem {
  id: string;
  title?: string;
  href?: string;
  disabled?: boolean;
  external?: boolean;
  icon?: keyof typeof Icons;
  label?: string;
  description?: string;
  active?: boolean;
}

export interface NavItemWithChildren extends NavItem {
  items: NavItemWithChildren[];
}
export interface NavItemWithOptionalChildren extends NavItem {
  items?: NavItemWithChildren[];
}

export type MainNavItem = NavItemWithOptionalChildren;
export type SidebarNavItem = NavItemWithChildren;


export type SiteConfig = typeof siteConfig;
export const siteConfig = {
  mainNav: [
    {
      id: "Articles",
      title: "Articles",
      items: [
        {
          id: "ArticleList",
          href: "/article/all",
          title: "All Articles",
          description: "All articles on this site",
          items: [],
        },
        {
          id: "ArticleTags",
          href: "/article/tags",
          title: "Tags",
          description: "All tags used in articles",
          items: [],
        },
      ],
    },
    {
      id: "About",
      title: "About",
      href: "/page/about",
    },
  ] satisfies MainNavItem[],
};