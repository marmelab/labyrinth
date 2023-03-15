<?php

namespace App\Controller;

use Doctrine\DBAL\Exception\UniqueConstraintViolationException;
use Doctrine\Persistence\ManagerRegistry;
use Symfony\Bundle\SecurityBundle\Security;
use Symfony\Component\PasswordHasher\Hasher\UserPasswordHasherInterface;
use Symfony\Component\Routing\Annotation\Route;
use Symfony\Component\HttpFoundation\Request;
use Symfony\Component\HttpFoundation\Response;
use Symfony\Component\Security\Http\Authentication\AuthenticationUtils;

use App\Entity\User;
use Symfony\Component\Form\FormError;

#[Route('/auth')]
class AuthController extends AuthBaseController
{
    function __construct(
        protected ManagerRegistry $doctrine,
        protected UserPasswordHasherInterface $passwordHasher,
    ) {
        parent::__construct($doctrine->getManager(), $passwordHasher);
    }

    #[Route('/sign-up', name: 'auth_sign_up', methods: 'GET')]
    public function getSignUp(): Response
    {
        $form = $this->createSignUpForm();
        return $this->render('auth/sign_up.html.twig', [
            'form' => $form,
        ]);
    }

    #[Route('/sign-up-post', name: 'auth_sign_up_post', methods: 'POST')]
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

    #[Route(path: '/sign-in', name: 'app_login')]
    public function login(AuthenticationUtils $authenticationUtils): Response
    {
        if ($this->getUser()) {
            return $this->redirectToRoute('home');
        }

        $error = $authenticationUtils->getLastAuthenticationError();
        $lastUsername = $authenticationUtils->getLastUsername();

        return $this->render('auth/sign_in.html.twig', ['last_username' => $lastUsername, 'error' => $error]);
    }

    #[Route(path: '/sign-out', name: 'auth_sign_out', methods: 'POST')]
    public function logout(Security $security): Response
    {
        if ($this->getUser()) {
            return $security->logout();
        }
        return $this->redirect('home');
    }
}
