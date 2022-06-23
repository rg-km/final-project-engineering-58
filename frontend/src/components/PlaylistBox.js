import React from 'react';
// import PlaylistItem from '../PlaylistItem';
import "../assets/css/PlaylistBox.css"
import "../assets/css/PlaylistItem.css"
// import PlaylistItemStyle from '../assets/styles/PlaylistItemStyle';
import { Link, NavLink } from 'react-router-dom'

const PlaylistItem = ({ video, active, played }) => {
  const activeStyle = ({isActive}) => {
    return {
      background: isActive ? '#CBCBCB' : 'none',
      borderRadius: isActive ? '50px' : '0px',
      textDecoration: 'none',
    }
  }

  return (
    <NavLink style={activeStyle} to={{ pathname: `/${video.id}`, autoplay: true}} className="container-item-playlist">
      <div className="video-number">{video.num}</div>
      <div className="video-title">{video.title}</div>
    </NavLink>
  );
}

const PlaylistBox = ({videos, active}) => (
  <div className="box-playlist">
    {videos.map(video => (
      <PlaylistItem
        key={video.id}
        video={video}
        active={video.id === active.id}
        played={video.played}
      />
    ))}
  </div>
)

export default PlaylistBox;