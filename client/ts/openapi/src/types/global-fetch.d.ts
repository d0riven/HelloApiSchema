// GlobalFetchがTS3.6.2から未対応になったのでマッピングし直している
// https://github.com/apollographql/apollo-link/issues/1131
declare type GlobalFetch = WindowOrWorkerGlobalScope;
