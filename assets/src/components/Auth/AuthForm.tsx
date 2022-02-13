import {Fragment, useContext, useEffect, useState} from 'react';
import {useNavigate} from 'react-router-dom';
import {
    Box,
    Button,
    Container,
    Flex,
    FormControl,
    FormLabel,
    Heading,
    HStack,
    Input,
    InputGroup,
    InputRightElement,
    Stack,
    useColorModeValue,
} from '@chakra-ui/react';
import {ViewIcon, ViewOffIcon} from '@chakra-ui/icons';
import AuthContext from '../../store/auth-context';
import Errors from '../Errors/Errors';

interface AuthFormProps {
    fromSignUp?: boolean;
}

const AuthForm = (props: AuthFormProps) => {
    const navigate = useNavigate();
    const [email, setEmail] = useState('');
    const [username, setUsername] = useState('');
    const [password, setPassword] = useState('');

    const authContext = useContext(AuthContext);

    const [loggingIn, setLoggingIn] = useState(!props.fromSignUp);
    const [errors, setErrors] = useState({});

    useEffect(() => {
        setLoggingIn(!props.fromSignUp ?? false);
    }, [props.fromSignUp])

    const switchModeHandler = () => {
        setLoggingIn((prevState) => !prevState);
        setErrors({});
    };

    const endpoint = loggingIn ? '/api/signin' : '/api/signup'

    async function submitHandler(event: React.FormEvent) {
        event.preventDefault();
        setErrors({});

        try {
            const response = await fetch(endpoint,
                {
                    method: 'POST',
                    body: JSON.stringify({
                        email: email,
                        Username: username,
                        Password: password,
                    }),
                    headers: {
                        'Content-Type': 'application/json',
                    },
                }
            );
            const data = await response.json();
            if (!response.ok) {
                let errorText = loggingIn ? 'Login failed' : 'Sign up failed';
                if (!data.hasOwnProperty('error')) {
                    console.log("this happen1");
                    throw new Error(errorText);
                }
                if ((typeof data['error'] === 'string')) {
                    console.log("this happen2");
                    setErrors({'unknown': data['error']})
                } else {
                    console.log("this happen3");
                    setErrors(data['error']);
                }
            } else {
                console.log("this happen4");
                authContext.login(data.jwt)
                navigate('/', {replace: true});
            }
        } catch (error) {
            console.log("this happen5");
            let message = 'Unknown Error'
            if (error instanceof Error) message = error.message
            setErrors({"error": message});
        }
    }

    const header = loggingIn ? 'Login' : 'Sign up';
    const mainButtonText = loggingIn ? 'Login' : 'Create account';
    const switchModeButtonText = loggingIn ? 'Create new account' : 'Login with existing account';
    const errorContent = Object.keys(errors).length === 0 ? null : Errors(errors);
    const [showPassword, setShowPassword] = useState(false);

    return (
        <Container>
            <Flex
                align={'center'}
                justify={'center'}
                bg={useColorModeValue('gray.50', 'gray.800')}>
                <Stack spacing={8} mx={'auto'} maxW={'lg'} py={12} px={6} minW={'lg'}>
                    <Stack align={'center'}>
                        <Heading fontSize={'4xl'} textAlign={'center'}>
                            {header}
                        </Heading>
                    </Stack>
                    <Box
                        rounded={'lg'}
                        bg={useColorModeValue('white', 'gray.700')}
                        boxShadow={'lg'}
                        p={8}>
                        <Stack spacing={4}>
                            <form onSubmit={submitHandler}>
                                <HStack security={'0'}>
                                    {loggingIn && (
                                        <FormControl id="username" isRequired>
                                            <FormLabel>Username</FormLabel>
                                            <Input
                                                value={username}
                                                onChange={(event) => setUsername(event.target.value)}
                                                type="text"
                                                placeholder={'Username'}/>
                                        </FormControl>
                                    )}
                                    {!loggingIn && (
                                        <Fragment>
                                            <Box>
                                                <FormControl id="username" isRequired>
                                                    <FormLabel>Username</FormLabel>
                                                    <Input
                                                        value={username}
                                                        onChange={(event) => setUsername(event.target.value)}
                                                        type="text"
                                                        placeholder={'Username'}/>
                                                </FormControl>
                                            </Box>
                                            <Box>
                                                <FormControl id="email" isRequired>
                                                    <FormLabel>Email address</FormLabel>
                                                    <Input
                                                        value={email}
                                                        onChange={(event) => setEmail(event.target.value)}
                                                        type="text"
                                                        placeholder={'Email'}/>
                                                </FormControl>
                                            </Box>
                                        </Fragment>
                                    )}
                                </HStack>
                                <FormControl id="password" isRequired>
                                    <FormLabel>Password</FormLabel>
                                    <InputGroup>
                                        <Input type={showPassword ? 'text' : 'password'} value={password}
                                               onChange={(event) => setPassword(event.target.value)}
                                               placeholder={'Password'}/>
                                        <InputRightElement h={'full'}>
                                            <Button
                                                variant={'ghost'}
                                                onClick={() =>
                                                    setShowPassword((showPassword) => !showPassword)
                                                }>
                                                {showPassword ? <ViewIcon/> : <ViewOffIcon/>}
                                            </Button>
                                        </InputRightElement>
                                    </InputGroup>
                                </FormControl>
                                <Stack spacing={10} pt={2}>
                                    <Button
                                        type="submit"
                                        loadingText="Submitting"
                                        size="lg"
                                        bg={'blue.400'}
                                        color={'white'}
                                        _hover={{
                                            bg: 'blue.500',
                                        }
                                        }>
                                        {mainButtonText}
                                    </Button>
                                </Stack>
                                <Stack pt={6}>
                                    <Button
                                        onClick={switchModeHandler}
                                        loadingText="Submitting"
                                        size="lg"
                                        bg={'green.400'}
                                        color={'white'}
                                        _hover={{
                                            bg: 'green.500',
                                        }
                                        }>
                                        {switchModeButtonText}
                                    </Button>
                                </Stack>
                            </form>
                            {errorContent}
                        </Stack>
                    </Box>
                </Stack>
            </Flex>
        </Container>
    );
}

export default AuthForm;
