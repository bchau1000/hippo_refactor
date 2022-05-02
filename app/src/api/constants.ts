/*
 * Define constants here as modules to make code readable
 * and easy to navigate
 */

module ServerInfo {
    const serverHost:string = "localhost";
    const serverPort:string = "3000"
    export const serverAddress:string = `http://${serverHost}:${serverPort}`;
}

module ContentType {
    export const json:HeadersInit = { 'Content-Type': 'application/json' };
}

export {
    ServerInfo,
    ContentType
}