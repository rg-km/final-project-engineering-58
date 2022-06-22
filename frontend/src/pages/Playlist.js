import React from 'react';
import PlaylistBox from '../components/PlaylistBox';
import "../assets/css/Playlist.css";

const PlaylistHeader = ({ active }) => {
  return (
    <div className="header-playlist">
      <p>{active.title}</p>
    </div>
  );
}

const Playlist = ({videos, active}) => {
return (
    <div className='container-playlist'>
      <PlaylistHeader active={active} total={videos.length} />
      <PlaylistBox videos={videos} active={active} />
    </div>
  );
}

export default Playlist;