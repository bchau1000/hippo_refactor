import { serverAddress } from "./constants";

const login = async (event: any): Promise<Response> => {
    event.preventDefault();

    const { email, password }:
        { email: string, password: string } = event.detail;

    const body: string = JSON.stringify({
        email: email,
        password: password,
    });

    const settings: RequestInit = {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: body,
    };

    const response: Response = await fetch(
        `${serverAddress}/api/login`,
        settings);

    return response;
}

export {
    login
}