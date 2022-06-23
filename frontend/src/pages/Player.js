import React, {useState, useEffect} from 'react';
import Video from '../components/Video';
import Playlist from './Playlist';
import "../assets/css/Player.css";
import VideoList from '../VideoList';
import NavVideo from '../components/NavVideo';
import { useNavigate, useParams, useLocation } from 'react-router-dom';

const Player = () => {
  const videoList = VideoList;
  const navigate = useNavigate();
  const location = useLocation();
  let params = useParams();

  const [state, setState] = useState({
    videos: videoList.playlist,
    activeVideo: videoList.playlist[0],
    playlistId: videoList.playlistId,
    autoplay: false
  });

  useEffect(() => {
    const currVideo = params.activeVideo;
    if(currVideo !== undefined){
      const newActiveVideo = state.videos.findIndex(
        video => video.id === currVideo
      )

      setState(prev => ({
        ...prev,
        activeVideo: prev.videos[newActiveVideo],
        autoplay: location.autoplay
      }));

    } else {
      navigate({
        pathname: `/${state.activeVideo.id}`,
        autoplay: false
      });
    }
  }, [navigate, location.autoplay, params.activeVideo, state.activeVideo.id, state.videos])

  const handleEndCallback = () => {
    const currVideo = params.activeVideo;
    const currVideoIndex = state.videos.findIndex(
      video => video.id === currVideo
    );
    const nextVideo = currVideoIndex === state.videos.length - 1 ? 0 : currVideoIndex + 1;

    if(nextVideo) {
      navigate({
        pathname: `/${state.videos[nextVideo].id}`,
        autoplay: false
      });
    } else {
      navigate({
        pathname: `/${state.activeVideo.id}`,
        autoplay: false
      });
    }
  }

  return (
    <div>
      <div>
      <NavVideo />
        {state.videos !== null ? (
          <div className="container-player">
            <Playlist 
              videos={state.videos}
              active={state.activeVideo}
            />
            <Video 
              active={state.activeVideo}
              autoplay={state.autoplay}
              endCallback={handleEndCallback}
            />
          </div>
        ) : null}
      </div>
    </div>
  )
}

export default Player;