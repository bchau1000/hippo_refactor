/**
 * Define constants here as modules to make code readable
 * and easy to navigate
 */
module ServerInfo {
    const serverHost:string = '127.0.0.1';
    const serverPort:string = '4000'
    export const serverAddress:string = `http://${serverHost}:${serverPort}`;
}

module ContentType {
    export const json:HeadersInit = { 'Content-Type': 'application/json' };
}

/**
 * Constants to label events to dispatch
 */
module Events {
    export const login:string = 'login';
    export const register:string = 'register';
}

export {
    ServerInfo,
    ContentType,
    Events
}