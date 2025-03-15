import React, { useRef, useEffect } from "react";
import videojs from "video.js";
import Player from "video.js/dist/types/player";
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

  return (
      <div className="flex-1 p-4">
        <div className="relative aspect-video w-full rounded-lg bg-black">
          <div className="flex w-full aspect-video">
            <video ref={videoRef} className="video-js w-full h-full" />
          </div>
        </div>
      </div>
  );
};