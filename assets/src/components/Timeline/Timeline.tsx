import { useState, useContext } from 'react';

import AuthContext from '../../store/auth-context';
import Errors from '../Errors/Errors';
import {Container, Heading, SimpleGrid, Stack} from "@chakra-ui/react";
import PostSubmitter from "../Posts/PostSubmitter";
import Post from "../Posts/Post";

interface TimelineProps {
    accountId?: number;
}

const Timeline = (props: TimelineProps) => {
    const [editing, setEditing] = useState(false);
    const [errors, setErrors] = useState({});
    const posts = [{}]

    const authContext = useContext(AuthContext);
    //const isLoggedIn = authContext.loggedIn;
    const isLoggedIn = true;

    const switchModeHandler = () => {
        setEditing((prevState) => !prevState);
        setErrors({});
    };

    async function deleteHandler() {

            const response = await fetch('api/posts/' + "someid",
                {
                    method: 'DELETE',
                    headers: {
                        'Authorization': 'Bearer ' + authContext.token,
                    },
                }
            );
            const data = await response.json();
            if (!response.ok) {
                let errorText = 'Failed to add new post.';
                if (!data.hasOwnProperty('error')) {
                    throw new Error(errorText);
                }
                if ((typeof data['error'] === 'string')) {
                    setErrors({ 'unknown': data['error'] })
                } else {
                    setErrors(data['error']);
                }
            } else {
                //props.onDeletePost(props.post.ID);
            }
    }

    const editPostHandler = () => {
        setEditing(false);
        //props.onEditPost();
    }

    const switchModeButtonText = editing ? 'Cancel' : 'Edit';
    const cardButtons = editing ?
        <div className="container">
            <button type="button" className="btn btn-link" onClick={switchModeHandler}>{switchModeButtonText}</button>
            <button type="button" className="btn btn-danger float-right mx-3" onClick={deleteHandler}>Delete</button>
        </div>
        :
        <div className="container">
            <button type="button" className="btn btn-link" onClick={switchModeHandler}>{switchModeButtonText}</button>
            <button type="button" className="btn btn-danger float-right mx-3" onClick={deleteHandler}>Delete</button>
        </div>
    const errorContent = Object.keys(errors).length === 0 ? null : Errors(errors);

    return (
        <Container py={12}>
                <Stack spacing={4}>
                    <Heading>Timeline placeholder</Heading>
                    {isLoggedIn && (<PostSubmitter/>)}
                    <Post name={"DingDong"} email={"danny-duller@hotmail.com"} message={"hey der wathopaedsopd"} date={new Date(2020, 9)}/>
                    <Post name={"DingDong"} email={"danny-duller@hotmail.com"} message={"hey der wathopaedsopd"} date={new Date(2020, 9)}/>
                    <Post name={"DingDong"} email={"danny-duller@hotmail.com"} message={"hey der wathopaedsopd"} date={new Date(2020, 9)}/>
                    <Post name={"DingDong"} email={"danny-duller@hotmail.com"} message={"hey der fdsdsfd sdfdsfwathopaedsopd"} date={new Date(2020, 9)}/>
                </Stack>
        </Container>
    );
};

export default Timeline;
