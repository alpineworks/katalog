"use client";

import ColorBadge from "@/components/colorbadge/colorbadge";
import { Button } from "@/components/ui/button";
import { ColumnDef } from "@tanstack/react-table";
import { ArrowUpDown } from "lucide-react";
import { formatDistance } from "date-fns";

// This type is used to define the shape of our data.
// You can use a Zod schema here if you want.
export type Deployment = {
  id: string;
  name: string;
  namespace: string;
  cluster: string;
  replicas: number;
  true_replicas: number;
  created_at: string;
  updated_at: string;
  labels: { [key: string]: string };
};

export type Deployments = {
  deployments: Deployment[];
};

export const columns: ColumnDef<Deployment>[] = [
  {
    accessorKey: "cluster",
    header: "Cluster",
    cell: ({ row }) => {
      return (
        <div>
          <ColorBadge text={row.original.cluster} />
        </div>
      );
    },
  },
  {
    accessorKey: "namespace",
    header: "Namespace",
    cell: ({ row }) => {
      return (
        <div>
          <ColorBadge text={row.original.namespace} />
        </div>
      );
    },
  },
  {
    accessorKey: "name",
    header: "Name",
  },
  {
    accessorKey: "replicas",
    header: ({ column }) => {
      return (
        <div className="-translate-x-4">
          <Button
            variant="ghost"
            onClick={() => column.toggleSorting(column.getIsSorted() === "asc")}
          >
            Replicas
            <ArrowUpDown className="ml-2 h-4 w-4" />
          </Button>
        </div>
      );
    },
  },
  {
    accessorKey: "true_replicas",
    header: "True Replicas",
  },
  {
    accessorKey: "labels",
    header: "Labels",
    cell: ({ row }) => {
      return (
        <div className="flex flex-wrap gap-2">
          {Object.entries(row.original.labels).map(([key, value]) => {
            const plaintext = `${key}:${value}`;
            return <ColorBadge text={plaintext} key={plaintext} />;
          })}
        </div>
      );
    },
  },
  {
    accessorKey: "created_at",
    header: "Created",
    cell: ({ row }) => {
      let stringifiedSince = formatDistance(
        new Date(row.original.created_at),
        new Date(),
        { addSuffix: true }
      );

      return <div className="w-32">{stringifiedSince}</div>;
    },
  },
  {
    accessorKey: "updated_at",
    header: "Updated",
    cell: ({ row }) => {
      let stringifiedSince = formatDistance(
        new Date(row.original.updated_at),
        new Date(),
        { addSuffix: true }
      );

      return <div className="w-32">{stringifiedSince}</div>;
    },
  },
];
