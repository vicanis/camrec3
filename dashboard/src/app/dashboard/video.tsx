"use client";

import useVisible from "@/hooks/usevisible";
import {
    Fragment,
    HTMLAttributes,
    ReactNode,
    useCallback,
    useEffect,
    useRef,
    useState,
} from "react";

export default function Video({ ts }: { ts: string }) {
    const ref = useRef<HTMLDivElement>(null);

    const isVisible = useVisible(ref);

    const [isPlaying, setIsPlaying] = useState(false);
    const [isLoading, setIsLoading] = useState(false);
    const [isError, setIsError] = useState(false);

    const clickHandler = useCallback(
        (player: HTMLVideoElement) => {
            if (!isPlaying) {
                player.play();
                setIsPlaying(true);
            } else {
                player.pause();
                setIsPlaying(false);
            }
        },
        [isPlaying]
    );

    const loadHandler = useCallback(() => {
        setIsLoading(false);
    }, []);

    useEffect(() => {
        if (isVisible) {
            setIsLoading(true);
        }
    }, [isVisible]);

    if (isError) {
        return <Placeholder>Error</Placeholder>;
    }

    return (
        <Fragment>
            {(!isVisible || isLoading) && (
                <div ref={ref}>
                    <Placeholder>Loading ...</Placeholder>
                </div>
            )}
            {isVisible && (
                <video
                    autoPlay
                    onClick={(event) =>
                        clickHandler(event.target as HTMLVideoElement)
                    }
                    onEnded={(event) =>
                        (event.target as HTMLVideoElement).play()
                    }
                    src={"/api/event/" + ts}
                    onLoadedData={loadHandler}
                    onError={(event) => {
                        console.error("load failed", event);
                        setIsError(true);
                    }}
                    style={{ display: isLoading ? "none" : "block" }}
                />
            )}
        </Fragment>
    );
}

function Placeholder({
    children,
    ...rest
}: { children: ReactNode } & HTMLAttributes<HTMLDivElement>) {
    return (
        <div
            className="max-w-3xl aspect-video bg-slate-300 flex items-center justify-center"
            {...rest}
        >
            {children}
        </div>
    );
}
