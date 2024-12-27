import { Article } from "@/api/types";
import { absoluteUrl, cn } from "@/lib/utils";
import React from "react";
import {
  Table,
  TableBody,
  TableCaption,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from "@/components/ui/table";
import { Button } from "../ui/button";
import Link from "next/link";

type Props = {
  articles?: Article[];
  tag_id?: number;
  className?: string;
};

function renderHeader() {
  return (
    <TableHeader className="sticky">
      <TableRow>
        <TableHead key="id" className="w-[50px] bg-stone-100">
          ID
        </TableHead>
        <TableHead key="title" className="w-[200px] bg-stone-100">
          Title
        </TableHead>
        <TableHead key="desc" className="w-[400px] bg-stone-100">
          Desc
        </TableHead>
        <TableHead key="content" className="w-[200px] bg-stone-100">
          Content
        </TableHead>
        <TableHead key="image_url" className="w-[200px] bg-stone-100">
          ImageURL
        </TableHead>
        <TableHead key="state" className="w-[50px] bg-stone-100">
          Stage
        </TableHead>
        <TableHead key="tag" className="w-[100px] bg-stone-100">
          Tag
        </TableHead>
      </TableRow>
    </TableHeader>
  );
}

function renderCell(articles: Article[]) {
  return (
    <TableBody className="h-96 overflow-y-auto">
      {articles?.map((article: Article, index: number) => {
        return (
          <TableRow key={`row-${index}`}>
            <TableCell>{article.id}</TableCell>
            <TableCell>{article.title}</TableCell>
            <TableCell>{article.desc}</TableCell>
            <TableCell>{article.content}</TableCell>
            <TableCell>{article.cover_image_url}</TableCell>
            <TableCell>{article.state}</TableCell>
            <TableCell className="flex-col items-center justify-center space-y-2">
              {article.tags?.map((tag) => {
                return (
                  <Link key={tag.id} href={absoluteUrl("article/all", new URLSearchParams({
                    tag_id: tag.id.toString()
                  }))}>
                    <Button>{tag.name}</Button>
                  </Link>
                );
              })}
            </TableCell>
          </TableRow>
        );
      })}
    </TableBody>
  );
}

export default function ArticleList({ articles, className, tag_id }: Props) {
  if (articles && articles.length > 0) {
    return (
      <Table
        className={cn(className, "my-1 border border-solid border-inherit")}
      >
        {!!tag_id && <TableCaption>Article of tag {tag_id}</TableCaption>}
        {renderHeader()}
        {renderCell(articles)}
      </Table>
    );
  }
  return <div>article-list</div>;
}
