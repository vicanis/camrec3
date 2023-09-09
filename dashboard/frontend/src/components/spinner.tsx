import Icon from "@mdi/react";
import { mdiLoading } from "@mdi/js";

export default function Spinner() {
    return (
        <div className={`animate-spin h-10 w-10 mx-auto`}>
            <Icon path={mdiLoading} size={1.5} />
        </div>
    );
}
