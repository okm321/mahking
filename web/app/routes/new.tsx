import { createFileRoute } from "@tanstack/react-router";
import { GroupCreateView } from "~/views/GroupCreateView";

export const Route = createFileRoute("/new")({
  head: () => ({
    meta: [
      { title: "新規作成 | Mahking 麻雀の対局記録アプリ" },
      { name: "description", content: "Welcome to React Router!" },
    ],
  }),
  component: NewPage,
});

function NewPage() {
  return <GroupCreateView />;
}
