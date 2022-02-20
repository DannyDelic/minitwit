import {useState, useContext, useEffect} from 'react';

import AuthContext from '../../store/auth-context';
import Errors from '../Errors/Errors';
import {Container, Heading, SimpleGrid, Stack} from "@chakra-ui/react";
import PostSubmitter from "../Posts/PostSubmitter";
import Post, {IPost} from "../Posts/Post";
import PostList from "../Posts/PostList";

interface TimelineProps {
    accountId?: number;
}

const Timeline = (props: TimelineProps) => {
    const [editing, setEditing] = useState(false);
    const [errors, setErrors] = useState({});
    const [posts, setPosts] = useState<IPost[]>([]);

    const authContext = useContext(AuthContext);
    const isLoggedIn = authContext.loggedIn;

    const switchModeHandler = () => {
        setEditing((prevState) => !prevState);
        setErrors({});
    };

    async function getAllPostsHandler() {

            const response = await fetch('api/msgs/',
                {
                    method: 'GET',
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
                setPosts(jsonPostsToPosts(data.data));
            }
    }

    useEffect(() => {
        getAllPostsHandler();
    }, []);

    const editPostHandler = () => {
        setEditing(false);
        //props.onEditPost();
    }

    const switchModeButtonText = editing ? 'Cancel' : 'Edit';
    const cardButtons = editing ?
        <div className="container">
            <button type="button" className="btn btn-link" onClick={switchModeHandler}>{switchModeButtonText}</button>
            <button type="button" className="btn btn-danger float-right mx-3" >Delete</button>
        </div>
        :
        <div className="container">
            <button type="button" className="btn btn-link" onClick={switchModeHandler}>{switchModeButtonText}</button>
            <button type="button" className="btn btn-danger float-right mx-3" >Delete</button>
        </div>
    const errorContent = Object.keys(errors).length === 0 ? null : Errors(errors);

    return (
        <Container py={12}>
                <Stack spacing={4}>
                    <Heading>Timeline placeholder</Heading>
                    {isLoggedIn && (<PostSubmitter/>)}
                    <PostList posts={posts}/>
                </Stack>
        </Container>
    );
};

const jsonPostsToPosts = (objects: [{[key: string]: string; }]) : IPost[] => {
    return objects.map((object: {[key: string]: string; }) => {
        let newObj: IPost = {content: "", name: "", date: new Date()};
        newObj.content = object['content'];
        newObj.name = object['poster'];
        newObj.date = new Date(Date.parse(object['created_at']));

        return newObj;
    });
}


export default Timeline;
