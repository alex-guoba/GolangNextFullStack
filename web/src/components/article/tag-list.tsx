import { Tag } from "@/api/types";
import { cn } from "@/lib/utils";
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

type Props = {
  tags?: Tag[];
  name?: string;
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
          Name
        </TableHead>
        <TableHead key="desc" className="w-[400px] bg-stone-100">
          State
        </TableHead>
        <TableHead key="content" className="w-[200px] bg-stone-100">
          CreateBy
        </TableHead>
      </TableRow>
    </TableHeader>
  );
}

function renderCell(tags: Tag[]) {
  return (
    <TableBody className="h-96 overflow-y-auto">
      {tags?.map((tag: Tag, index: number) => {
        return (
          <TableRow key={`row-${index}`}>
            <TableCell>{tag.id}</TableCell>
            <TableCell>{tag.name}</TableCell>
            <TableCell>{tag.state}</TableCell>
            <TableCell>{tag.created_by}</TableCell>
          </TableRow>
        );
      })}
    </TableBody>
  );
}

export default function TagList({ tags, className, name }: Props) {
  if (tags && tags.length > 0) {
    return (
      <Table
        className={cn(className, "my-1 border border-solid border-inherit")}
      >
        {!!name && <TableCaption>Tag of {name}</TableCaption>}
        {renderHeader()}
        {renderCell(tags)}
      </Table>
    );
  }
  return <div>article-list</div>;
}
