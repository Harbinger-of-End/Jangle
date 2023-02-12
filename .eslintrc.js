module.exports = {
    extends: [
        'plugin:@next/eslint-plugin-next/recommended',
        'prettier:recommended',
        'plugin:typescript/recommended',
    ],
    parserOptions: {
        ecmaVersion: 'latest',
        sourceType: 'module',
        ecmaFeatures: {
            jsx: true,
        },
        jsx: true,
    },
    rules: {
        indent: [4, 'error'],
    },
};
