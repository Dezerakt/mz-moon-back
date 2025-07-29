'use client';

import AudioPlayer from 'react-h5-audio-player';
import 'react-h5-audio-player/lib/styles.css';

export default function Player() {
    return <div className="p-4">
        <AudioPlayer
            src="http://localhost:8080/v1/song/1"
            autoPlay={false}
        />
    </div>;
}