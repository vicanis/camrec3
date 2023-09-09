import { ReactNode } from "react";
import Spinner from "./spinner";

export default function Loader({ children }: { children?: ReactNode }) {
    return (
        <div className="h-full w-full flex items-center justify-center">
            <div className="bg-white py-6 px-10 shadow-xl rounded-md">
                <Spinner />
                {typeof children !== "undefined" && (
                    <div className="pt-4">{children}</div>
                )}
            </div>
        </div>
    );
}
