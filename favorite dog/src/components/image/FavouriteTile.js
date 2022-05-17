import React from "react";
import styled from "styled-components";
import { Col, ResponsiveEmbed } from "react-bootstrap";

const StyledCol = styled(Col) `
flex: 1 0 30%;
margin: 10px;
padding-top: 10px;
padding-bottom: 10px;
align-items: center;
background-color: #fff;
border-radius: 5px;
box-shadow: 0 5px 15px rgba(0, 0, 0, 0.3);
display: flex;
`;

const StyledDiv = styled.div`
    padding: 10px;
`;

const StyledSpan = styled.span`
    color: red;
    padding: 10px;
`;

function FavouriteTile(props) {
	return (
        <StyledCol md={4}>
            <ResponsiveEmbed>
                <embed src = {props.imageUrl} alt = "Error in loading"/>
            </ResponsiveEmbed>
		</StyledCol>
	);
}

export default FavouriteTile;