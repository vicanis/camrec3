import { Outlet } from "react-router-dom";
import Menu from "../components/menu";

export default function PageIndex() {
    return (
        <div className="p-3">
            <Menu />
            <hr />
            <Outlet />
        </div>
    );
}
