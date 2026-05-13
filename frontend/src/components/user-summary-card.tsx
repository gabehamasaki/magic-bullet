import { ShieldCheck, UserRound } from "lucide-react";

import { Button } from "@/components/ui/button";

type UserSummaryCardProps = {
    user: {
        id: string;
        email: string;
        role: string;
    } | null;
    errorMessage: string;
};

export function UserSummaryCard({ user, errorMessage }: UserSummaryCardProps) {
    return (
        <section className="rounded-[28px] border border-border bg-card p-6 shadow-2xl shadow-black/8 md:p-8">
            <div className="flex flex-col gap-8">
                <div className="flex items-start justify-between gap-4">
                    <div className="space-y-3">
                        <div className="inline-flex items-center gap-2 rounded-full border border-border bg-secondary px-3 py-1 text-xs font-medium uppercase tracking-[0.24em] text-muted-foreground">
                            <span className="size-2 rounded-full bg-primary" />
                            Astro + React + shadcn/ui
                        </div>
                        <div className="space-y-2">
                            <h1 className="text-4xl font-semibold tracking-tight text-foreground md:text-5xl">
                                Magic Bullet
                            </h1>
                            <p className="max-w-xl text-sm leading-6 text-muted-foreground md:text-base">
                                Boilerplate com OpenAPI tipado de ponta a ponta, Astro SSR no frontend e componentes preparados para React com shadcn/ui.
                            </p>
                        </div>
                    </div>
                    <div className="hidden rounded-3xl border border-border bg-secondary p-4 md:block">
                        <UserRound className="size-8 text-primary" />
                    </div>
                </div>

                {user ? (
          <div className="grid gap-4 md:grid-cols-3">
            <div className="min-w-0 rounded-2xl border border-border bg-secondary p-4">
              <p className="text-xs uppercase tracking-[0.2em] text-muted-foreground">ID</p>
              <p className="mt-3 break-words text-sm font-medium text-foreground">{user.id}</p>
            </div>
            <div className="min-w-0 rounded-2xl border border-border bg-secondary p-4">
              <p className="text-xs uppercase tracking-[0.2em] text-muted-foreground">Email</p>
              <p className="mt-3 break-all text-sm font-medium text-foreground">{user.email}</p>
            </div>
            <div className="min-w-0 rounded-2xl border border-border bg-secondary p-4">
              <p className="text-xs uppercase tracking-[0.2em] text-muted-foreground">Role</p>
              <p className="mt-3 inline-flex max-w-full items-center gap-2 break-words text-sm font-medium text-foreground">
                <ShieldCheck className="size-4 text-primary" />
                {user.role}
              </p>
                        </div>
                    </div>
                ) : (
                    <div className="rounded-2xl border border-destructive/20 bg-destructive/10 p-4 text-sm text-destructive">
                        {errorMessage || "No user payload returned."}
                    </div>
                )}

                <div className="flex flex-wrap gap-3">
                    <Button>shadcn/ui ready</Button>
                    <Button variant="outline">OpenAPI typed client</Button>
                </div>
            </div>
        </section>
    );
}
