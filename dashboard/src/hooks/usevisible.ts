import { RefObject, useEffect, useState } from "react";

export default function useVisible(ref: RefObject<HTMLElement>) {
    const [visible, setVisible] = useState(false);

    useEffect(() => {
        const target = ref.current;

        if (target === null) {
            return;
        }

        const observer = new IntersectionObserver((entries) => {
            for (const entry of entries) {
                if (entry.isIntersecting) {
                    setVisible(true);
                }
            }
        });

        observer.observe(target);

        return () => observer.disconnect();
    }, [ref]);

    return visible;
}
