import {
    createUserWithEmailAndPassword,
    signInWithEmailAndPassword,
    onAuthStateChanged,
    getAuth
} from "firebase/auth";

import {auth} from './firebase';
import type { UserCredential, Auth } from "firebase/auth"

const login = async (event: any): Promise<string> => {
    event.preventDefault();
    const { email, password }:
        { email: string, password: string } = event.detail;
    try {
        const userCredential: UserCredential = await signInWithEmailAndPassword(auth, email, password);
        console.log(userCredential);
    } catch (error: any) {
        const errorCode = error.code;
        const errorMessage = error.message;
        console.error(`${errorCode}: ${errorMessage}`);
    }

    return await Promise.resolve("");
}

const register = async (event: any): Promise<void> => {
    event.preventDefault();
    const auth:Auth = getAuth();
    const { displayName, email, password }:
        { displayName: string, email: string, password: string } = event.detail;

    try {
        const userCredential: UserCredential = await createUserWithEmailAndPassword(
            auth,
            email,
            password
        );
        console.log(userCredential.user);
    } catch (error: any) {
        const errorCode = error.code;
        const errorMessage = error.message;

        console.error(`${errorCode}: ${errorMessage}`);
    }
}

const authenticate = ():string => {
    let uid = "";

    onAuthStateChanged(auth, (user) => {
        if (user) {
            console.log(user);
            uid = user.uid;
        }
    });

    return uid;
}

export {
    login,
    register,
    authenticate
}