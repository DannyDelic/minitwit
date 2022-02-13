import {Fragment} from 'react';
import NavigationBar from "./NavigationBar";
import {Flex, useColorModeValue} from "@chakra-ui/react";

const Layout = (props: { children: React.ReactNode }) => {
    return (
        <Fragment>
            <NavigationBar/>
            <main>
                <Flex minH={'100vh'} bg={useColorModeValue('gray.50', 'gray.50')}>
                    {props.children}
                </Flex>
            </main>
        </Fragment>
    );
};

export default Layout;
