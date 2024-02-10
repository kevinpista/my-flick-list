import YouTube from 'react-youtube';
import Modal from 'react-modal';

const YouTubeModal = ({ isOpen, setIsOpen, videoId }) => { // setIsOpen is function from parent component to close or open modal
    const customStyles = {
        overlay: {
            zIndex: 1000, 
            backgroundColor: 'rgba(0,0,0,0.88)',
            display: 'flex',
            alignItems: 'center',
            justifyContent: 'center',
            position: 'fixed',
        },
        content: {
            position: 'absolute',
            left: 'auto',
            right: 'auto',
            bottom: 'auto',
            top: '50%',
            transform: 'translateY(-50%)',
            backgroundColor: 'black',
            outline: 'none',
            border: 'none',
            paddingTop: '10px',
            paddingBottom: '10px',
            paddingRight: '0px',
            paddingLeft: '0px',
        },
    };

    const opts = {
      height: '546rem',
      width: '896rem',
      playerVars: {
        autoplay: 1,
      },
    };

    // const handleOpen = () => setIsOpen(true); // this will be invoked by the parent function calling this component
    const handleClose = () => {
        setIsOpen(false); // trigger the callback to change the isOpen in the parent component
    }    
    return (
      <div>
        <Modal 
            isOpen={isOpen}
            onOverlayClick={handleClose} // Closes modal if user clicks outside
            onRequestClose={handleClose} // Closes modal on 'ESC' key
            style={customStyles}
            contentLabel="YouTube Modal"
        >
          <YouTube 
            videoId={videoId} 
            opts={opts} 
          />
        </Modal>
      </div>
    );
  };
  export default YouTubeModal;