import React from 'react';
import s from "./Scan.module.css";
import { useNavigate  } from "react-router-dom";

const Scan = () => {
    const formRef = React.useRef(null);
    const inputRef = React.useRef(null);
    const labelRef = React.useRef(null);
    const navigate = useNavigate();

    var personsResponseJson;

    const formClick = async () => {
        const formData = new FormData();
        
        formData.append('file', inputRef.current.files[0]);
        
        const options = {
          method: 'POST',
          body: formData,
        };
        labelRef.current.innerText = "Подождите..."
        var personsResponse = await fetch('http://localhost:49812/emotion/recognize', options);
        personsResponseJson = await personsResponse.json()

        navigate("/results",{state: {personsResponseJson}});
    }

  return (
    <div>

      
      <form action="http://localhost:49812/emotion/recognize" method="post" id="upload-form" encType="multipart/form-data" ref={formRef} >

        <input
          type="file"
          name="file"
          id="file"
          className="file inputfile-2"
          data-multiple-caption="{count} files selected"
          multiple=""
          onChange={formClick}
          ref={inputRef}
        />
        <label htmlFor="file" ref={labelRef}>
          <img src="./img/upload.png" width="53" height="51"></img>
        </label>
      </form>
    </div>
  );
};

export default Scan;
