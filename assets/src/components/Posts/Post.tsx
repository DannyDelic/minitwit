import {
    Box,
    Image,
    Text,
    chakra,
    Flex,
    SimpleGrid,
    Stat,
    StatLabel,
    useColorModeValue,
} from '@chakra-ui/react';
import CryptoJS from "crypto-js/x64-core";
import MD5 from "crypto-js/hmac-md5";

interface PostProps {
    name: string;
    email: string
    message: string;
    date: Date;
}

function Post(props: PostProps) {
    const { name, message, date, email } = props;
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
                <Box
                    my={'auto'}
                    color={useColorModeValue('gray.800', 'gray.200')}
                    alignContent={'center'}>
                    <Image src={getGravatarUrl(email, 48)}/>
                </Box>
                <Box pl={{ base: 2, md: 4 }}>
                    <Text fontWeight={'medium'} display={"flex"} >
                        {name}
                        <Text color={'gray'} fontWeight={"light"} paddingLeft={1}>{`— ${dateString} @ ${time}`}</Text>
                    </Text>
                    <Text fontSize={'2xl'} fontWeight={'medium'}>
                        {message}
                    </Text>
                </Box>
            </Flex>
        </Stat>
    );
}

function getGravatarUrl(email: string, size: number=80) {
    const hex = MD5(email.trim().toLowerCase(), "").toString(CryptoJS.enc.Hex);
    return `http://www.gravatar.com/avatar/${hex}?d=identicon&s=${size}`;
}

export default Post;
