import { ReactNode } from "react";

export default function Blur({
    children,
    mode = "fixed",
    ...props
}: {
    children?: ReactNode;
    mode?: "fixed" | "absolute";
} & React.HTMLAttributes<HTMLDivElement>) {
    return (
        <div
            className={`${
                mode === "fixed" ? "fixed" : "absolute"
            } top-0 left-0 w-full h-full backdrop-blur-md z-10`}
            {...props}
        >
            {children}
        </div>
    );
}
