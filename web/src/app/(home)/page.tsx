import Shell from "@/components/shell/shell";
import {
  PageHeader,
  PageHeaderHeading,
  PageHeaderDescription,
} from "@/components/page-header";

export default function Home() {
  return (
    <Shell className="md:pb-10">
      <PageHeader>
        <PageHeaderHeading>Home</PageHeaderHeading>
        <PageHeaderDescription>Home page description</PageHeaderDescription>
      </PageHeader>
    </Shell>
  );
}
