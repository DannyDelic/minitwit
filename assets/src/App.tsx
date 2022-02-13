import { Routes, Navigate, Route} from 'react-router-dom';

import Layout from './components/Layout/Layout';
import {Fragment, useContext} from 'react';
import AuthContext from './store/auth-context';
import {ChakraProvider, Flex, useColorModeValue} from "@chakra-ui/react";
import * as React from "react";
import Timeline from "./components/Timeline/Timeline";
import AuthPage from "./pages/AuthPage";
import TimelinePage from "./pages/Timeline";

function App() {
    const authContext = useContext(AuthContext);

    return (
        <ChakraProvider>
            <Layout>
                <Routes>
                    <Route path="/" element={<TimelinePage/>}/>
                    <Route path="/public" element={<TimelinePage/>}/>
                    {!authContext.loggedIn && (
                        <Fragment>
                            <Route path="/login" element={<AuthPage fromSignUp={false}/>}/>
                            <Route path="/signup" element={<AuthPage fromSignUp={true}/>}/>
                        </Fragment>
                    )}
                    {/*<Route path="/posts">
                        {authContext.loggedIn && <PostsPage />}
                        {!authContext.loggedIn && <Redirect to="/auth" />}
                    </Route>*/}
                    <Route path="*" element={<Navigate to="/"/>}/>
                </Routes>
            </Layout>
        </ChakraProvider>
    )
}

export default App;
