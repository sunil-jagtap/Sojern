import React, { useEffect } from "react";
import styled from "styled-components";
import { useAppContext } from "../../AppContext";
import GallerySet from './GallerySet';
import { Button } from "react-bootstrap";
import useHttp from '../../hooks/http';
import { useHistory } from "react-router";
import SpinnerWithOverlay from '../spinner/SpinnerWithOverlay';

const TitleContainer = styled.div`
	display: flex;
	justify-content: space-between;
	padding: 20px 0px 24px 0px;
	background-color: #fff!important;
	border-radius: 5px!important;
    box-shadow: 0 2px 10px rgba(0, 0, 0, 0.3)!important;
	margin-top: 10px!important;
    margin-bottom: 10px!important;
`;

const Title = styled.h1`
	&&& {
		font-family: Helvetica Neue;
		font-size: 24px;
		font-weight: 400;
		line-height: inherit;
	}
`;

const MainContainer = styled.div`
	background: rgb(255, 255, 255);
	bottom: 17px;
	right: 21px;
	left: 208px;
	top: 146px;
`;

const StyledSpan = styled.span`
	margin-top: 20px;
	float: right;
	margin-right: 15px;
`;

const StyledButton = styled(Button)`
    margin: 5px;
`;

function GalleryPage() {
    const { state, dispatch } = useAppContext();
    const { getRandomImages } = useHttp();
    const { loading } = state;
    const history = useHistory();

    useEffect(() => {
		    getRandomImages(dispatch);
	}, []);

	return (
		<>
            <TitleContainer>
                <span>
					<iframe src="https://giphy.com/embed/X9FTTSAIz4g3rvCFK9" width="300" height="100" frameBorder="0" class="giphy-embed" ></iframe><p><a href="https://giphy.com/gifs/hallmarkecards-puppy-hallmark-ecards-X9FTTSAIz4g3rvCFK9"></a></p>
                </span>
                <StyledSpan>
                    <StyledButton onClick = { () => {history.push("/favourites")}}>Favourites</StyledButton>
                    <Button onClick = {() => getRandomImages(dispatch)}>Refresh</Button>
                </StyledSpan>
			</TitleContainer>
            <MainContainer>
                <GallerySet />
                <SpinnerWithOverlay loading = { loading } spinnerLabel = {"loading"}/>
            </MainContainer>
		</>
	);
}

export default GalleryPage;
