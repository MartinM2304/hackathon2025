import React, { useRef, useEffect, use } from "react";
import videojs from "video.js";
import Player from "video.js/dist/types/player";
import "video.js/dist/video-js.css";
import "video.js/dist/video-js.css";

interface VideoPlayerProps {
  src: string;
}

export const VideoPlayer: React.FC<VideoPlayerProps> = ({ src }) => {
  const videoRef = useRef<HTMLVideoElement | null>(null);
  const playerRef = useRef<Player | null>(null);

  useEffect(() => {
    if (!videoRef.current) return;

    playerRef.current = videojs(videoRef.current, {
      controls: true,
      autoplay: false,
      responsive: true,
      sources: [{ src, type: "application/x-mpegURL" }],
    });

    return () => {
      if (playerRef.current) {
        playerRef.current.dispose();
        playerRef.current = null;
      }
    };
  }, [src]);
  
  useEffect(() => {
    const player = playerRef.current;

    return () => {
        if(player && !player.isDisposed()) {
            player.dispose();
            playerRef.current = null;
        }
    };
  }, [src]);

  return (
      <div className="flex w-full md:p-4 justify-center items-center">
        <div className="relative aspect-video w-full rounded-lg bg-black">
          <div className="flex w-full aspect-video">
            <video ref={videoRef} className="video-js w-full h-full" />
          </div>
        </div>
      </div>
  );
};