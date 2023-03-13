<?php

namespace App\Controller;

use Doctrine\DBAL\Exception\UniqueConstraintViolationException;
use Doctrine\Persistence\ManagerRegistry;
use Symfony\Component\PasswordHasher\Hasher\UserPasswordHasherInterface;
use Symfony\Component\Routing\Annotation\Route;
use Symfony\Component\HttpFoundation\Request;
use Symfony\Component\HttpFoundation\Response;

use App\Entity\User;
use Symfony\Component\Form\FormError;

#[Route('/auth', name: 'auth_')]
class AuthController extends AuthBaseController
{
    function __construct(
        protected ManagerRegistry $doctrine,
        protected UserPasswordHasherInterface $passwordHasher,
    ) {
        parent::__construct($doctrine->getManager(), $passwordHasher);
    }

    #[Route('/sign-up', name: 'sign_up', methods: 'GET')]
    public function getSignUp(): Response
    {
        $form = $this->createSignUpForm();
        return $this->render('auth/sign_up.html.twig', [
            'form' => $form,
        ]);
    }

    #[Route('/sign-up-post', name: 'sign_up_post', methods: 'POST')]
    public function postSignUp(Request $request): Response
    {
        $user = new User();
        $form = $this->createSignUpForm($user)->handleRequest($request);
        if (!$form->isSubmitted() || !$form->isValid()) {
            return $this->render('auth/sign_up.html.twig', [
                'form' => $form,
            ]);
        }

        try {
            $this->signUpUser($user);
            return $this->redirectToRoute('home');
        } catch (UniqueConstraintViolationException $e) {
            $form->addError(new FormError("Username or E-mail is already used by another user."));
            return $this->render('auth/sign_up.html.twig', [
                'form' => $form,
            ]);
        }
    }
}
