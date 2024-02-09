import React, { useState, useEffect } from 'react';
import YouTube from 'react-youtube';
import Modal from 'react-modal';

const YouTubeModal = ({ isOpen, videoId  }) => {
    const [modalOpen, setModalOpen] = useState(isOpen); 

    const customStyles = {
        overlay: {zIndex: 1000} // Makes sure the modal is always on top
      };

    useEffect(() => {
        setModalOpen(isOpen);
    }, [isOpen]);

    const opts = {
      height: '390',
      width: '640',
      playerVars: {

      },
    };

    // const handleOpen = () => setIsOpen(true); // this will be invoked by the parent function calling this component
    const handleClose = () => setModalOpen(false);
    
    return (
      <div>
        <Modal 
            isOpen={modalOpen} 
            onClose={handleClose} 
            closeTimeoutMS={200} 
            style = {customStyles}
        >
          <h2>Trailer</h2>
          <YouTube videoId={videoId} opts={opts} />
        </Modal>
      </div>
    );
  };
  
  export default YouTubeModal;