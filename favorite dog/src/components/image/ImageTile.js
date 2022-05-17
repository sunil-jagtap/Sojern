import React, { useState } from "react";
import styled from "styled-components";
import { useAppContext } from "../../AppContext";
import { Col, Image, ResponsiveEmbed } from "react-bootstrap";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { faHeart } from "@fortawesome/free-solid-svg-icons";
import { faHeart as farHeart } from "@fortawesome/free-regular-svg-icons";
import useLocalStorage from "../../hooks/localstorage";


const StyledCol = styled(Col) `
    flex: 1 0 30%;
    margin: 10px;
    padding-top: 10px;
    padding-bottom: 10px;
    align-items: center;
    background-color: #fff;
	border-radius: 5px;
    box-shadow: 0 2px 10px rgba(0, 0, 0, 0.3);
    display: flex;
`;

const StyledDiv = styled.div`
    padding: 10px;
`;

const StyledSpan = styled.span`
    color: red;
    padding: 10px;
`;

function ImageTile(props) {
    const { state, dispatch } = useAppContext();
    const { addToLocalStorage } = useLocalStorage();
    const [isAddedToFav, setIsAddedToFav] = useState(false);

    const addToFavourite = (imageUrl) => {
        setIsAddedToFav(true);
        addToLocalStorage(imageUrl);
    };

	return (
        <StyledCol md={4}>
            <ResponsiveEmbed>
                <embed src = {props.imageUrl} alt = "Error in loading"/>
            </ResponsiveEmbed>
            <StyledDiv>
                {
                   isAddedToFav ?  
                   <>
                   <FontAwesomeIcon
					style={{
						color: "red",
						marginTop: "12px",
						marginRight: "12px"
                    }}
					icon={faHeart}/>
                    <StyledSpan>Added</StyledSpan>
                    </> :
                    <> 
                    <FontAwesomeIcon
					style={{
						color: "red",
						marginTop: "12px",
						marginRight: "12px"
                    }}
                    onClick = { () => addToFavourite(props.imageUrl) }
					icon={farHeart}/>
                    <StyledSpan>Add to Favourites</StyledSpan>
                    </>
                }
            </StyledDiv>
		</StyledCol>
	);
}

export default ImageTile;