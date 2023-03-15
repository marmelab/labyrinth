<?php

namespace App\Controller;

use Doctrine\DBAL\Exception\UniqueConstraintViolationException;
use Doctrine\Persistence\ManagerRegistry;
use Symfony\Bundle\SecurityBundle\Security;
use Symfony\Component\HttpFoundation\JsonResponse;
use Symfony\Component\HttpFoundation\Request;
use Symfony\Component\PasswordHasher\Hasher\UserPasswordHasherInterface;
use Symfony\Component\Routing\Annotation\Route;
use Symfony\Component\Serializer\SerializerInterface;
use Symfony\Component\Validator\Validator\ValidatorInterface;

use App\Entity\User;
use App\Security\TokenServiceInterface;

#[Route('/api/v1/auth')]
class AuthApiController extends AuthBaseController
{
    public function __construct(
        protected ManagerRegistry $doctrine,
        protected UserPasswordHasherInterface $passwordHasher,
        protected SerializerInterface $serializer,
        protected ValidatorInterface $validator,
        protected TokenServiceInterface $tokenService,
    ) {
        parent::__construct($doctrine->getManager(), $passwordHasher, $tokenService);
    }

    #[Route('/identity', name: 'auth_api_identity', methods: 'GET')]
    public function getIdentity(): JsonResponse
    {
        return $this->json([
            'data' => $this->getUser(),
        ]);
    }

    #[Route('/sign-up', name: 'auth_api_sign_up_post', methods: 'POST')]
    public function postSignUp(Request $request): JsonResponse
    {
        $user = $this->serializer->deserialize($request->getContent(),  User::class, 'json');
        $errors = $this->validator->validate($user);
        if ($errors->count() > 0) {
            return $this->json([
                'error' => $errors->__toString(),
            ], 400);
        }

        try {
            return $this->json([
                'data' => $this->signUpUser($user),
            ]);
        } catch (UniqueConstraintViolationException) {
            return $this->json([
                'error' => 'Username or E-mail is already used by another user.',
            ], 400);
        }
    }

    #[Route('/sign-out', name: 'auth_api_sign_out', methods: 'POST')]
    public function postSignOut(Security $security): JsonResponse
    {
        $user = $this->getUser();
        if ($user) {
            $security->logout(false);
        }
        return $this->json([
            'data' => null,
        ]);
    }

    #[Route(path: '/sign-in', name: 'api_login', methods: 'POST')]
    public function login(): JsonResponse
    {
        /** @var User */
        $user = $this->getUser();
        if ($user) {
            $reponse = $this->json([
                'data' => $user,
            ]);
            return $this->tokenService->setToken($reponse, $user);
        }

        return $this->json([
            'data' => null,
        ]);
    }

    #[Route(path: '/check', methods: 'GET')]
    public function check(): JsonResponse
    {
        /** @var User */
        $user = $this->getUser();
        if ($user) {
            $reponse = $this->json([
                'data' => $user,
            ]);
            return $this->tokenService->setToken($reponse, $user);
        }

        return $this->json([
            'data' => null,
        ], 401);
    }
}
