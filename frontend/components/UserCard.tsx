'use client';

import Link from 'next/link';
import { User } from '@/lib/api/users';
import FollowButton from './FollowButton';

interface UserCardProps {
    user: User;
    showFollow?: boolean;
}

export default function UserCard({ user, showFollow = true }: UserCardProps) {
    return (
        <div className="flex items-center justify-between p-3 bg-white/5 border border-white/10 rounded-xl hover:border-white/20 transition-all">
            <Link href={`/profile/${user.id}`} className="flex items-center gap-3 group">
                {user.avatar_url ? (
                    <img
                        src={user.avatar_url}
                        alt={user.username}
                        className="w-10 h-10 rounded-full object-cover ring-2 ring-purple-500/30 group-hover:ring-purple-500/60 transition-all"
                    />
                ) : (
                    <div className="w-10 h-10 rounded-full bg-gradient-to-br from-purple-500 to-cyan-500 flex items-center justify-center text-white font-bold text-sm ring-2 ring-purple-500/30">
                        {user.username.charAt(0).toUpperCase()}
                    </div>
                )}
                <div>
                    <p className="text-sm font-semibold text-white group-hover:text-purple-300 transition-colors">
                        {user.username}
                    </p>
                    <p className="text-xs text-slate-500">{user.phone}</p>
                </div>
            </Link>
            {showFollow && <FollowButton targetId={user.id} />}
        </div>
    );
}
