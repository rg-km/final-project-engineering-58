import React from 'react';
// import PlaylistItem from '../PlaylistItem';
import "../assets/css/PlaylistBox.css"
import "../assets/css/PlaylistItem.css"
// import PlaylistItemStyle from '../assets/styles/PlaylistItemStyle';
import { Link, NavLink } from 'react-router-dom'

const PlaylistItem = ({ video, active, played }) => {
  const activeStyle = ({isActive}) => {
    return {
      fontWeight: isActive ? '600' : '',
      borderRadius: isActive ? '30px' : '0px',
    }
  }

  return (
    <div className="container-item-playlist">
        <div className="video-number">{video.num}</div>
        <div className="video-title">
        {/* <NavLink style={activeStyle}> */}
          <NavLink style={activeStyle} to={{ pathname: `/${video.id}`, autoplay: true}}>
              {video.title}
          </NavLink></div>
          {/* </NavLink> */}
    </div>
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