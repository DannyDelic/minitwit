import {
    Box,
    Image,
    Flex,
    Stat,
    useColorModeValue, HStack, Text
} from '@chakra-ui/react';
import CryptoJS from "crypto-js/x64-core";
import MD5 from "crypto-js/hmac-md5";

interface PostProps {
    name: string;
    message: string;
    date: Date;
}

export interface IPost {
    name: string;
    content: string;
    date: Date;
}

function Post(props: PostProps) {
    const { name, message, date } = props;
    const dateString = date.toLocaleDateString()
    const time = date.toLocaleTimeString()


    return (
        <Stat
            px={{ base: 2, md: 4 }}
            py={'5'}
            shadow={'xl'}
            bg={useColorModeValue('white', 'whiteAlpha.100')}
            rounded={'lg'}>
            <Flex>
                <HStack security={'0'}>
                <Box
                    my={'auto'}
                    color={useColorModeValue('gray.800', 'gray.200')}
                    alignContent={'center'}
                    minW={16}
                    maxW={16}>
                    <Image src={getGravatarUrl(name, 48)}/>
                </Box>
                <Box pl={{ base: 2, md: 4 }}>
                    <HStack security={'0'}>
                        <Text fontWeight={'medium'} display={"flex"}>{name}</Text>
                        <Text color={'gray'} fontWeight={"light"} paddingLeft={1}>{`— ${dateString} @ ${time}`}</Text>
                    </HStack >
                    <Text fontSize={'2xl'} overflowWrap={"break-word"} wordBreak={"break-word"}>
                        {message}
                    </Text>
                </Box>
                </HStack>
            </Flex>
        </Stat>
    );
}

function getGravatarUrl(name: string, size: number=80) {
    const hex = MD5(name.trim().toLowerCase(), "").toString(CryptoJS.enc.Hex);
    return `http://www.gravatar.com/avatar/${hex}?d=identicon&s=${size}`;
}

export default Post;
