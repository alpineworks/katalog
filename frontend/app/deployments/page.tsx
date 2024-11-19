import { Session } from "next-auth";
import { redirect } from "next/navigation";
import { auth } from "@/lib/auth/auth";
import { DataTable } from "./data-table";
import { columns, Deployments } from "./columns";

async function getDeployments(session: Session): Promise<Deployments> {
  const res = await fetch(
    `${process.env.BACKEND_HOST_URL}/api/v1/deployments?email=${session.user.email}&provider=${session.user.provider}`,
    {
      method: "GET",
      headers: {
        "Content-Type": "application/json",
        Authorization: `Bearer ${session.token}`,
      },
    }
  );
  if (res.status === 200) {
    return res.json() as Promise<Deployments>;
  } else {
    throw new Error("failed to get deployments");
  }
}

export default async function DeploymentsPage() {
  const session = await auth();
  if (!session) {
    redirect("/unauthorized");
  }

  const deployments = await getDeployments(session);

  return (
    <div className="flex justify-center min-h-screen pt-20">
      <div className="container">
        <h1 className="text-4xl font-semibold text-left">Deployments</h1>
        <DataTable columns={columns} data={deployments.deployments} />
      </div>
    </div>
  );
}
