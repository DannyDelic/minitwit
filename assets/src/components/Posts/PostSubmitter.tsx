import { FormEvent, ChangeEvent, useState } from 'react';
import {
    Stack,
    Input,
    Button,
    useColorModeValue,
    Heading,
    Text,
    Container,
    Flex, FormControl,
} from '@chakra-ui/react';
import { CheckIcon } from '@chakra-ui/icons';

function PostSubmitter() {
    const [email, setEmail] = useState('');
    const [state, setState] = useState<'initial' | 'submitting' | 'success'>(
        'initial'
    );
    const [error, setError] = useState(false);

    return (
        <Flex
            align={'center'}
            justify={'center'}>
            <Container
                maxW={'lg'}
                bg={useColorModeValue('white', 'whiteAlpha.100')}
                boxShadow={'xl'}
                rounded={'lg'}
                p={6}>
                <Heading
                    as={'h2'}
                    fontSize={{ base: 'xl', sm: '2xl' }}
                    textAlign={'center'}
                    mb={5}>
                    What's on your mind?
                </Heading>
                <Stack
                    direction={{ base: 'column', md: 'row' }}
                    as={'form'}
                    spacing={'12px'}
                    onSubmit={(e: FormEvent) => {
                        e.preventDefault();
                        setError(false);
                        setState('submitting');

                        // remove this code and implement your submit logic right here
                        setTimeout(() => {
                            if (email === 'fail@example.com') {
                                setError(true);
                                setState('initial');
                                return;
                            }

                            setState('success');
                        }, 1000);
                    }}>
                    <FormControl>
                        <Input
                            variant={'solid'}
                            borderWidth={1}
                            color={'gray.800'}
                            _placeholder={{
                                color: 'gray.400',
                            }}
                            borderColor={useColorModeValue('gray.300', 'gray.700')}
                            id={'email'}
                            type={'text'}
                            required
                            placeholder={'Who asked, just close the browser'}
                            aria-label={'Your post'}
                            value={email}
                            disabled={state !== 'initial'}
                            onChange={(e: ChangeEvent<HTMLInputElement>) =>
                                setEmail(e.target.value)
                            }
                        />
                    </FormControl>
                    <FormControl w={{ base: '100%', md: '40%' }}>
                        <Button
                            colorScheme={state === 'success' ? 'green' : 'blue'}
                            isLoading={state === 'submitting'}
                            w="100%"
                            type={state === 'success' ? 'button' : 'submit'}>
                            {state === 'success' ? <CheckIcon /> : 'Share'}
                        </Button>
                    </FormControl>
                </Stack>
                <Text
                    mt={2}
                    textAlign={'center'}
                    color={error ? 'red.500' : 'gray.500'}>
                    {error && 'Oh no an error occured! ???? Please try again later.'}
                </Text>
            </Container>
        </Flex>
    );
}

export default PostSubmitter;