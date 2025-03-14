import React, {useRef, useEffect} from "react";
import videojs from "video.js";
import Player from "video.js/dist/types/player";


interface VideoPlayerProps {
    src: string;
}

export const VideoPlayer: React.FC<VideoPlayerProps> = ({src}) => {
    const videoRef = useRef<HTMLVideoElement | null>(null);
    const playerRef = useRef<Player | null>(null);


    useEffect(() => {
        if(!playerRef.current && videoRef.current) {
            const videoElement = document.createElement("video-js");

            videoRef.current.appendChild(videoElement);

            playerRef.current = videojs(videoElement, {
                controls: true,
                fluid: true,
                sources: [{ src, type: "application/x-mpegURL" }],
        
            });
        }
    }, [src]);

    useEffect(() => {
        const player = playerRef.current;
        
        return () => {
            if(player && !player.isDisposed()) {
                player.dispose();
                playerRef.current = null;
            }
        };
    }, []);

    return (
        <div data-vjs-player>
            <video ref={videoRef} />
        </div>
    );
}