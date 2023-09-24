import { ReactNode } from "react";

export default function DashboardLayout({ children }: { children: ReactNode }) {
    return <div className="max-w-3xl mx-auto py-3">{children}</div>;
}
