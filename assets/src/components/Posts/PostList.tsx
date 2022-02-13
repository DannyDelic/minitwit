import Post, {IPost} from "./Post";
import {Fragment} from "react";
import {Link} from "react-router-dom";

interface PostListProps {
    posts: IPost[];
}

function PostList(props: PostListProps) {
    const {posts} = props;


    return (
        <Fragment>
            {posts.map((post) => (
                <Post name={post.name} date={post.date} message={post.content}/>
            ))}
        </Fragment>
    );
}

export default PostList;
