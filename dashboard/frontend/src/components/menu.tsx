import { useLocation } from "react-router-dom";

const items: MenuItemType[] = [
    {
        name: "Главная",
        path: "/",
    },
    {
        name: "События",
        path: "/events",
    },
    {
        name: "Просмотр в реальном времени",
        path: "/live",
    },
];

export default function Menu() {
    return (
        <ul className="flex gap-4">
            {items.map((item) => (
                <MenuItem key={item.path} {...item} />
            ))}
        </ul>
    );
}

function MenuItem({ name, path }: MenuItemType) {
    const { pathname } = useLocation();

    return (
        <li>{pathname === path ? <b>{name}</b> : <a href={path}>{name}</a>}</li>
    );
}

type MenuItemType = {
    name: string;
    path: string;
};
