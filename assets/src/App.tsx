import { Routes, Navigate, Route} from 'react-router-dom';

import Layout from './components/Layout/Layout';
import { useContext } from 'react';
import AuthContext from './store/auth-context';
import {ChakraProvider, Flex, useColorModeValue} from "@chakra-ui/react";
import * as React from "react";
import Timeline from "./components/Timeline/Timeline";

function App() {
    const authContext = useContext(AuthContext);

    return (
        <ChakraProvider>
            <Layout>
                <Flex minH={'100vh'} bg={useColorModeValue('gray.50', 'gray.50')}>
                    <Timeline/>
                </Flex>
            </Layout>
        </ChakraProvider>
    )
}

export default App;
